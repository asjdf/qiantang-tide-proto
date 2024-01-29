/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type VersionRequest = {
}

export type VersionResponse = {
  version?: string
}

export class CommonService {
  static Version(req: VersionRequest, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<VersionRequest, VersionResponse>(`/gapi/common/v1/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}