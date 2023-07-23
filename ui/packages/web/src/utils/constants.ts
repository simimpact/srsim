// NOTE: othi: ping me on discord if remote api is out of date/500/404s

import { MvpWrapper } from "@/bindings/MvpWrapper";
import { AvatarConfig, EquipmentConfig } from "@/routes/Root/CharacterLineup";

export const OTHI_API = "https://api.othiremote.synology.me";
// export const OTHI_API = "http://127.0.0.1:5005/";

const API = {
  // WARN: :id does not actually mean id number, in this case it's string params (`danheng`)
  character: route<AvatarConfig>("/honkai/character/search/:id", "GET"),
  lightCone: route<EquipmentConfig>("/honkai/light_cone/search/:id", "GET"),
  mockHsrStat: route<MvpWrapper>("/utils/mock_hsr_stat", "GET"),
};

interface ApiRoute {
  path: string;
}
interface ApiGet<TResponse> {
  get: (params?: string | number) => Promise<TResponse>;
}
interface ApiPost<TPayload, TResponse> {
  post: (opt?: { payload?: TPayload; params?: string }) => Promise<TResponse>;
}

type Get<TRes> = ApiRoute & ApiGet<TRes>;
type Post<TReq, TRes> = ApiRoute & ApiPost<TReq, TRes>;
type GetPost<TReq, TRes> = ApiRoute & ApiGet<TRes> & ApiPost<TReq, TRes>;

type Method = "GET" | "POST" | undefined;
function route<TReq, TRes>(path: string): GetPost<TReq, TRes>;
function route<TRes>(path: string, method: "GET"): Get<TRes>;
function route<TReq, TRes>(path: string, method: "POST"): Post<TReq, TRes>;
function route<TReq, TRes>(
  path: string,
  method?: Method
): Get<TRes> | Post<TReq, TRes> | GetPost<TReq, TRes> {
  switch (method) {
    case "GET":
      return {
        path,
        get: async (params?: string | number) =>
          await serverFetch<TReq, TRes>(path, undefined, params),
      };
    case "POST":
      return {
        path,
        post: async (opt?: { payload?: TReq; params?: string }) =>
          await serverFetch<TReq, TRes>(
            path,
            { payload: opt?.payload, method: "POST" },
            opt?.params
          ),
      };
    default:
      return {
        // no method provided, allow both post and fetch
        path,
        get: async (params?: string | number) =>
          await serverFetch<TReq, TRes>(path, undefined, params),
        post: async (opt?: { payload?: TReq; params?: string }) =>
          await serverFetch<TReq, TRes>(
            path,
            { payload: opt?.payload, method: "POST" },
            opt?.params
          ),
      };
  }
}

export async function serverFetch<TPayload, TResponse>(
  endpoint: string,
  opt?: {
    payload?: TPayload;
    method: "POST" | "DELETE";
  },
  params?: string | number
): Promise<TResponse> {
  let url = OTHI_API + endpoint;
  if (params) {
    if (url.includes(":id")) url = url.replace(":id", `${params}`);
    else url += `/${params}`;
  }

  // POST
  if (opt) {
    const { payload, method } = opt;
    const body = JSON.stringify(payload);
    const res = await fetch(url, {
      body,
      headers: {
        "Content-Type": "application/json",
      },
      method,
    });

    if (res.ok) {
      return res.json() as Promise<TResponse>;
    } else {
      console.error("api fetch failed, code:", res.status);
      console.error("url:", url);
      const errText = await res.text();
      console.error("unknown error", errText);
      return Promise.reject(`unknown error ${errText}`);
    }
  } else {
    // GET
    const res = await fetch(url, {
      headers: {
        "Content-Type": "application/json",
      },
      method: "GET",
    });

    if (res.ok) {
      return res.json() as Promise<TResponse>;
    } else {
      console.error("api fetch failed, code:", res.status);
      console.error("url:", url);
      return Promise.reject(`unknown error ${await res.text()}`);
    }
  }
}

export default API;
