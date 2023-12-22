/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"
import * as TideV1Tide from "./tide.pb"
export type GetDayPredictTideRequest = {
  date?: GoogleProtobufTimestamp.Timestamp
}

export type GetDayPredictTideResponse = {
  tideList?: TideV1Tide.Tide[]
}

export type GetLocationPredictTideRequest = {
  locationId?: string
  date?: GoogleProtobufTimestamp.Timestamp
}

export type GetLocationPredictTideResponse = {
  tideList?: TideV1Tide.Tide[]
}

export type GetLocationListRequest = {
}

export type GetLocationListResponse = {
  locationList?: TideV1Tide.Location[]
}

export class TideService {
  static GetDayPredictTide(req: GetDayPredictTideRequest, initReq?: fm.InitReq): Promise<GetDayPredictTideResponse> {
    return fm.fetchReq<GetDayPredictTideRequest, GetDayPredictTideResponse>(`/gapi/tide/v1/predict/day?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetLocationPredictTide(req: GetLocationPredictTideRequest, initReq?: fm.InitReq): Promise<GetLocationPredictTideResponse> {
    return fm.fetchReq<GetLocationPredictTideRequest, GetLocationPredictTideResponse>(`/gapi/tide/v1/predict/location?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetLocationList(req: GetLocationListRequest, initReq?: fm.InitReq): Promise<GetLocationListResponse> {
    return fm.fetchReq<GetLocationListRequest, GetLocationListResponse>(`/gapi/tide/v1/location/list?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}