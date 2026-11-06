#!/usr/bin/env python3
"""Alat bantu PRD (hanya membaca inventaris; tidak mengubah apa pun).

  python3 -I prdtool.py check                  validasi alokasi modules.json, hitung per modul, laporkan typo/unassigned
  python3 -I prdtool.py names <id>             daftar nama fitur jatah modul <id> (mis. 06)
  python3 -I prdtool.py extract <id>           cetak Markdown detail semua fitur jatah modul <id>
  python3 -I prdtool.py coverage [<id> ...]    bandingkan docs/prd/NN-*.md dengan alokasi (nama fitur, ID unik, tautan)
"""
import json, os, re, sys

HERE = os.path.dirname(os.path.abspath(__file__))
PRD = os.path.dirname(HERE)


def norm(s):
    return re.sub(r"[^a-z0-9]+", " ", str(s).lower()).strip()


def load():
    cfg = json.load(open(os.path.join(HERE, "modules.json"), encoding="utf-8"))
    inv = {}
    for src, fn in cfg["sources"].items():
        data = json.load(open(os.path.join(HERE, fn), encoding="utf-8"))
        for d in data["result"]["inventory"]:
            inv[d["domain"]] = {"src": src, "title": d.get("title", d["domain"]), "features": d["features"]}
    return cfg, inv


def allocate(cfg, inv):
    """Kembalikan (per_module: id -> [(domain, feature)], errors, unassigned)."""
    errors, pick_idx, whole_idx = [], {}, {}
    for m in cfg["modules"]:
        for dom in m.get("whole", []):
            if dom in whole_idx:
                errors.append(f"domain {dom} dipakai whole oleh {whole_idx[dom]} dan {m['id']}")
            whole_idx[dom] = m["id"]
            if dom not in inv:
                errors.append(f"modul {m['id']}: domain whole tidak ada: {dom}")
        for dom, names in m.get("pick", {}).items():
            if dom not in inv:
                errors.append(f"modul {m['id']}: domain pick tidak ada: {dom}")
                continue
            have = {norm(f["name"]) for f in inv[dom]["features"]}
            for n in names:
                k = (dom, norm(n))
                if norm(n) not in have:
                    errors.append(f"modul {m['id']}: pick tidak ditemukan di {dom}: {n!r}")
                elif k in pick_idx:
                    errors.append(f"pick ganda {dom}/{n!r}: {pick_idx[k]} dan {m['id']}")
                pick_idx[k] = m["id"]
    per = {m["id"]: [] for m in cfg["modules"]}
    unassigned = []
    for dom, d in inv.items():
        for f in d["features"]:
            mid = pick_idx.get((dom, norm(f["name"]))) or whole_idx.get(dom)
            if mid is None:
                unassigned.append((dom, f["name"]))
            else:
                per[mid].append((dom, f))
    return per, errors, unassigned


def flagged(f):
    blob = (f.get("notes") or "") + " " + " ".join(f.get("subfeatures") or [])
    return "perlu verifikasi" in blob.lower()


def cmd_check(cfg, inv):
    per, errors, unassigned = allocate(cfg, inv)
    for e in errors:
        print("ERROR:", e)
    for dom, name in unassigned:
        print("UNASSIGNED:", dom, "|", name)
    tot_f = tot_s = 0
    for m in cfg["modules"]:
        fs = per[m["id"]]
        subs = sum(len(f.get("subfeatures") or []) for _, f in fs)
        doms = sorted({d for d, _ in fs})
        tot_f += len(fs); tot_s += subs
        print(f"{m['id']} {m['file']:<55} fitur={len(fs):4d} sub={subs:5d} ⚠={sum(flagged(f) for _, f in fs):3d}  <- {', '.join(doms)}")
    print(f"TOTAL fitur={tot_f} sub={tot_s} errors={len(errors)} unassigned={len(unassigned)}")
    return 1 if (errors or unassigned) else 0


def module(cfg, mid):
    for m in cfg["modules"]:
        if m["id"] == mid:
            return m
    sys.exit(f"modul tidak dikenal: {mid}")


def cmd_names(cfg, inv, mid):
    per, _, _ = allocate(cfg, inv)
    for dom, f in per[mid]:
        print(f"- [{dom}] {f['name']}  ({f.get('priority')}, {f.get('tier')})")


def cmd_extract(cfg, inv, mid):
    per, _, _ = allocate(cfg, inv)
    m = module(cfg, mid)
    print(f"# Jatah modul {m['id']} — {m['title']} (file {m['file']}, prefix {m['prefix']})\n")
    print(f"Jumlah fitur: {len(per[mid])}\n")
    cur = None
    for dom, f in per[mid]:
        if dom != cur:
            cur = dom
            print(f"\n## Sumber domain: {dom} — {inv[dom]['title']} ({inv[dom]['src']})\n")
        warn = " ⚠" if flagged(f) else ""
        print(f"### {f['name']}{warn}")
        print(f"- priority: {f.get('priority')} | tier: {f.get('tier')} | actors: {', '.join(f.get('actors') or [])}")
        print(f"- description: {f.get('description','')}")
        if f.get("notes"):
            print(f"- notes: {f['notes']}")
        subs = f.get("subfeatures") or []
        print(f"- subfeatures ({len(subs)}):")
        for i, s in enumerate(subs, 1):
            print(f"  {i}. {s}")
        print()


HEAD_RE = re.compile(r"^###\s+([A-Z][A-Z0-9-]*-\d{3})\s+(.+?)\s*$", re.M)
LINK_RE = re.compile(r"\[[^\]]*\]\(([^)#\s]+)(#[^)]*)?\)")


def cmd_coverage(cfg, inv, ids):
    per, _, _ = allocate(cfg, inv)
    mods = [m for m in cfg["modules"] if not ids or m["id"] in ids]
    all_ids, problems = {}, 0
    for m in mods:
        path = os.path.join(PRD, m["file"])
        fs = per[m["id"]]
        if not os.path.exists(path):
            print(f"{m['id']} {m['file']}: FILE TIDAK ADA ({len(fs)} fitur belum ditulis)"); problems += 1
            continue
        text = open(path, encoding="utf-8").read()
        heads = HEAD_RE.findall(text)
        titles = {norm(t) for _, t in heads}
        for al in re.findall(r"\*\*Alias inventaris:\*\*\s*(.+)", text):
            titles.update(norm(x) for x in al.split(";") if x.strip())
        missing = [f["name"] for _, f in fs if norm(f["name"]) not in titles]
        bad_prefix = [i for i, _ in heads if not i.startswith(m["prefix"] + "-")]
        for i, _ in heads:
            all_ids.setdefault(i, []).append(m["file"])
        rules = len(re.findall(r"^\s*\d+\.\s", text, re.M))
        subs = sum(len(f.get("subfeatures") or []) for _, f in fs)
        broken = []
        for target, _ in LINK_RE.findall(text):
            if target.startswith(("http://", "https://", "mailto:")):
                continue
            if not os.path.exists(os.path.normpath(os.path.join(os.path.dirname(path), target))):
                broken.append(target)
        tail_ok = "di luar cakupan" in text[-4000:].lower()
        status = "OK" if not missing and not bad_prefix and not broken and tail_ok else "PERLU PERBAIKAN"
        print(f"{m['id']} {m['file']}: {status} | heading={len(heads)}/{len(fs)} | butir bernomor={rules} vs sub-fitur={subs} ({(rules/subs*100) if subs else 0:.0f}%) | baris={text.count(chr(10))}")
        if missing:
            problems += 1
            print("   HILANG:", " ; ".join(missing))
        if bad_prefix:
            problems += 1
            print("   PREFIX SALAH:", ", ".join(bad_prefix))
        if broken:
            problems += 1
            print("   TAUTAN RUSAK:", ", ".join(sorted(set(broken))))
        if not tail_ok:
            problems += 1
            print("   TIDAK ADA bagian 'Di luar cakupan' di akhir file (file terpotong?)")
    dups = {i: f for i, f in all_ids.items() if len(f) > 1}
    for i, files in dups.items():
        problems += 1
        print("ID GANDA:", i, "di", ", ".join(files))
    print(f"SELESAI: {problems} masalah")
    return 1 if problems else 0


if __name__ == "__main__":
    a = sys.argv[1:]
    cfg, inv = load()
    if not a or a[0] == "check":
        sys.exit(cmd_check(cfg, inv))
    if a[0] == "names":
        cmd_names(cfg, inv, a[1])
    elif a[0] == "extract":
        cmd_extract(cfg, inv, a[1])
    elif a[0] == "coverage":
        sys.exit(cmd_coverage(cfg, inv, a[1:]))
    else:
        sys.exit(__doc__)
