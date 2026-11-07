/* eslint-disable */
/* tslint:disable */
// @ts-nocheck
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface DtoAuthLoginRequest {
  email?: string;
  password?: string;
}

export interface DtoAuthLoginResponse {
  token?: string;
  user?: DtoAuthRegisterResponse;
}

export interface DtoAuthRegisterRequest {
  email?: string;
  name?: string;
  password?: string;
  username?: string;
}

export interface DtoAuthRegisterResponse {
  created_at?: string;
  email?: string;
  id?: string;
  is_active?: boolean;
  name?: string;
  role?: string;
  username?: string;
}

export interface DtoUserUpdateProfileRequest {
  name?: string;
}

export interface DtoUserUserListResponse {
  pagination?: PkgPaginationResponse;
  users?: DtoUserUserResponse[];
}

export interface DtoUserUserResponse {
  created_at?: string;
  email?: string;
  id?: string;
  is_active?: boolean;
  name?: string;
  role?: string;
  updated_at?: string;
  username?: string;
}

export interface PkgPaginationResponse {
  page?: number;
  per_page?: number;
  total_data?: number;
  total_page?: number;
}

export interface ResponseResponse {
  data?: any;
  error?: boolean;
  errors?: any;
  message?: string;
  status?: number;
}

import type {
  AxiosInstance,
  AxiosRequestConfig,
  AxiosResponse,
  HeadersDefaults,
  ResponseType,
} from "axios";
import axios from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams
  extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<
  FullRequestParams,
  "body" | "method" | "query" | "path"
>;

export interface ApiConfig<SecurityDataType = unknown>
  extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  JsonApi = "application/vnd.api+json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({
    securityWorker,
    secure,
    format,
    ...axiosConfig
  }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({
      ...axiosConfig,
      baseURL: axiosConfig.baseURL || "//localhost:8080/api/v1",
    });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected mergeRequestParams(
    params1: AxiosRequestConfig,
    params2?: AxiosRequestConfig,
  ): AxiosRequestConfig {
    const method = params1.method || (params2 && params2.method);

    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...((method &&
          this.instance.defaults.headers[
            method.toLowerCase() as keyof HeadersDefaults
          ]) ||
          {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected stringifyFormItem(formItem: unknown) {
    if (typeof formItem === "object" && formItem !== null) {
      return JSON.stringify(formItem);
    } else {
      return `${formItem}`;
    }
  }

  protected createFormData(input: Record<string, unknown>): FormData {
    if (input instanceof FormData) {
      return input;
    }
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      const propertyContent: any[] =
        property instanceof Array ? property : [property];

      for (const formItem of propertyContent) {
        const isFileType = formItem instanceof Blob || formItem instanceof File;
        formData.append(
          key,
          isFileType ? formItem : this.stringifyFormItem(formItem),
        );
      }

      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = format || this.format || undefined;

    if (
      type === ContentType.FormData &&
      body &&
      body !== null &&
      typeof body === "object"
    ) {
      body = this.createFormData(body as Record<string, unknown>);
    }

    if (
      type === ContentType.Text &&
      body &&
      body !== null &&
      typeof body !== "string"
    ) {
      body = JSON.stringify(body);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(requestParams.headers || {}),
        ...(type ? { "Content-Type": type } : {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title Jari API
 * @version 1.0
 * @license MIT (https://opensource.org/licenses/MIT)
 * @termsOfService http://swagger.io/terms/
 * @baseUrl //localhost:8080/api/v1
 * @contact API Support <support@swagger.io> (http://www.swagger.io/support)
 *
 * REST API for Jari, an open-source work and project management platform (Jira clone). "Jari" is the working name until the final product name is decided.
 */
export class Api<
  SecurityDataType extends unknown,
> extends HttpClient<SecurityDataType> {
  api = {
    /**
     * @description Authenticate user with email and password
     *
     * @tags auth
     * @name V1AuthLoginCreate
     * @summary Login user
     * @request POST:/api/v1/auth/login
     */
    v1AuthLoginCreate: (
      request: DtoAuthLoginRequest,
      params: RequestParams = {},
    ) =>
      this.request<
        ResponseResponse & {
          data?: DtoAuthLoginResponse;
        },
        ResponseResponse
      >({
        path: `/api/v1/auth/login`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Register a new user with email, username, password, first name, and last name
     *
     * @tags auth
     * @name V1AuthRegisterCreate
     * @summary Register a new user
     * @request POST:/api/v1/auth/register
     */
    v1AuthRegisterCreate: (
      request: DtoAuthRegisterRequest,
      params: RequestParams = {},
    ) =>
      this.request<
        ResponseResponse & {
          data?: DtoAuthRegisterResponse;
        },
        ResponseResponse
      >({
        path: `/api/v1/auth/register`,
        method: "POST",
        body: request,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Get a paginated and filtered list of users. Admin only (temporary bridge until the permissions & roles module).
     *
     * @tags users
     * @name V1UsersList
     * @summary Get users list
     * @request GET:/api/v1/users
     * @secure
     */
    v1UsersList: (
      query?: {
        /**
         * Items per page
         * @default 10
         */
        per_page?: number;
        /**
         * Page number
         * @default 1
         */
        page?: number;
        /** Filter by user ID */
        id?: string;
        /** Filter by email */
        email?: string;
        /** Filter by username */
        username?: string;
        /** Filter by name (LIKE search) */
        name?: string;
        /** Filter by role */
        role?: string;
        /** Filter by multiple roles (comma-separated) */
        roles?: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<
        ResponseResponse & {
          data?: DtoUserUserListResponse;
        },
        ResponseResponse
      >({
        path: `/api/v1/users`,
        method: "GET",
        query: query,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Get the authenticated user's profile information
     *
     * @tags users
     * @name V1UsersProfileList
     * @summary Get user profile
     * @request GET:/api/v1/users/profile
     * @secure
     */
    v1UsersProfileList: (params: RequestParams = {}) =>
      this.request<
        ResponseResponse & {
          data?: DtoUserUserResponse;
        },
        ResponseResponse
      >({
        path: `/api/v1/users/profile`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Update the authenticated user's profile information
     *
     * @tags users
     * @name V1UsersProfileUpdate
     * @summary Update user profile
     * @request PUT:/api/v1/users/profile
     * @secure
     */
    v1UsersProfileUpdate: (
      request: DtoUserUpdateProfileRequest,
      params: RequestParams = {},
    ) =>
      this.request<
        ResponseResponse & {
          data?: DtoUserUserResponse;
        },
        ResponseResponse
      >({
        path: `/api/v1/users/profile`,
        method: "PUT",
        body: request,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
