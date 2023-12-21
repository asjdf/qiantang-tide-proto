/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"
export type Location = {
  unionId?: string
  name?: string
  latitude?: number
  longitude?: number
}

export type Tide = {
  id?: string
  locationId?: string
  time?: GoogleProtobufTimestamp.Timestamp
  tideHeight?: number
  surgeHeight?: number
  ornamentalLevel?: string
}