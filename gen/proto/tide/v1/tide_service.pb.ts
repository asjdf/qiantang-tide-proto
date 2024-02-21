/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"
import * as TideV1Tide from "./tide.pb"
export type GetDayPredictTideRequest = {
  date?: Date
}

export type GetDayPredictTideResponse = {
  tideList?: TideV1Tide.Tide[]
}

export type GetLocationPredictTideRequest = {
  locationId?: string
  date?: Date
}

export type GetLocationPredictTideResponse = {
  tideList?: TideV1Tide.Tide[]
}

export type GetLocationListRequest = {
}

export type GetLocationListResponse = {
  locationList?: TideV1Tide.Location[]
}

export type GetRealtimeTideRequest = {
}

export type GetRealtimeTideResponse = {
  realtimeTideList?: TideV1Tide.RealtimeTide[]
}

export type GetRealtimeTideOfLocationRequest = {
  locationId?: string
}

export type GetRealtimeTideOfLocationResponse = {
  realtimeTide?: TideV1Tide.RealtimeTide
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
  static GetRealtimeTide(req: GetRealtimeTideRequest, initReq?: fm.InitReq): Promise<GetRealtimeTideResponse> {
    return fm.fetchReq<GetRealtimeTideRequest, GetRealtimeTideResponse>(`/gapi/tide/v1/realtime?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetRealtimeTideOfLocation(req: GetRealtimeTideOfLocationRequest, initReq?: fm.InitReq): Promise<GetRealtimeTideOfLocationResponse> {
    return fm.fetchReq<GetRealtimeTideOfLocationRequest, GetRealtimeTideOfLocationResponse>(`/gapi/tide/v1/location/realtime?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}