package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mocks "app/internal/mocks/repository"
	"app/internal/shared/domain/entity"
	pkgjwt "app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

const roleTestSecret = "role-test-secret-value"

// newRoleRouter wires the same chain as the real route: language -> auth -> role -> handler.
func newRoleRouter(repo *mocks.MockUserRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(LanguageMiddleware())
	r.GET("/users", AuthMiddleware(roleTestSecret), RequireRole(repo, entity.RoleAdmin), func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}

func bearer(t *testing.T, userID string) string {
	t.Helper()
	token, err := pkgjwt.GenerateToken(roleTestSecret, pkgjwt.UserPayload{ID: userID, Email: "u@example.com", Username: "u"})
	assert.NoError(t, err)
	return "Bearer " + token
}

func getUsers(r *gin.Engine, authorization, lang string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}
	if lang != "" {
		req.Header.Set("Accept-Language", lang)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRequireRole_AdminAllowed(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(&entity.User{ID: "u1", Role: entity.RoleAdmin, IsActive: true}, nil)

	w := getUsers(newRoleRouter(repo), bearer(t, "u1"), "")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestRequireRole_RegularUserForbidden(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(&entity.User{ID: "u1", Role: entity.RoleUser, IsActive: true}, nil)

	w := getUsers(newRoleRouter(repo), bearer(t, "u1"), "")

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), `"message":"forbidden"`)
}

func TestRequireRole_ForbiddenMessageIsLocalized(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(&entity.User{ID: "u1", Role: entity.RoleUser, IsActive: true}, nil)

	w := getUsers(newRoleRouter(repo), bearer(t, "u1"), "id")

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "anda tidak memiliki izin untuk aksi ini")
}

func TestRequireRole_InactiveAdminForbidden(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(&entity.User{ID: "u1", Role: entity.RoleAdmin, IsActive: false}, nil)

	w := getUsers(newRoleRouter(repo), bearer(t, "u1"), "")

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRequireRole_NoTokenUnauthorized(t *testing.T) {
	repo := mocks.NewMockUserRepository(t) // no GetByID call expected

	w := getUsers(newRoleRouter(repo), "", "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireRole_MissingClaimsUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := mocks.NewMockUserRepository(t)
	r := gin.New()
	r.GET("/users", RequireRole(repo, entity.RoleAdmin), func(c *gin.Context) { c.String(http.StatusOK, "ok") })

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/users", nil))

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireRole_UserNotFoundUnauthorized(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "ghost").Return(nil, gorm.ErrRecordNotFound)

	w := getUsers(newRoleRouter(repo), bearer(t, "ghost"), "")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRequireRole_RepositoryErrorIs500(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(nil, errors.New("db down"))

	w := getUsers(newRoleRouter(repo), bearer(t, "u1"), "")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// A role change in the database applies to the very next request with the same token.
func TestRequireRole_RoleChangeAppliesWithoutNewToken(t *testing.T) {
	repo := mocks.NewMockUserRepository(t)
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(&entity.User{ID: "u1", Role: entity.RoleUser, IsActive: true}, nil).Once()
	repo.EXPECT().GetByID(mock.Anything, "u1").Return(&entity.User{ID: "u1", Role: entity.RoleAdmin, IsActive: true}, nil).Once()
	r := newRoleRouter(repo)
	token := bearer(t, "u1")

	assert.Equal(t, http.StatusForbidden, getUsers(r, token, "").Code)
	assert.Equal(t, http.StatusOK, getUsers(r, token, "").Code)
}
