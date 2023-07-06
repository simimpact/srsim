// NOTE: othi: ping me on discord if remote api is out of date/500/404s

// export const OTHI_API = "https://api.othiremote.synology.me/";
export const OTHI_API = "http://127.0.0.1:5005/";

export const ENDPOINT = {
  statMock: "utils/mock_hsr_stat",
  logMock: "utils/mock_hsr_log",
} as const;
export type Endpoint = (typeof ENDPOINT)[keyof typeof ENDPOINT];

// TODO: own file/hook
export async function typedFetch<TPayload, TResponse>(
  endpoint: (typeof ENDPOINT)[keyof typeof ENDPOINT],
  opt?: {
    payload?: TPayload;
    params?: string | number;
    method: "POST" | "DELETE";
  }
): Promise<TResponse> {
  let url = OTHI_API + endpoint;
  if (opt?.params) url += `/${opt.params}`;

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
      const err = await res.text();
      return Promise.reject(`unknown error\n${err}`);
    }
  }
}
