/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

export type Location = {
  unionId?: string
  name?: string
  latitude?: number
  longitude?: number
}

export type Tide = {
  id?: string
  locationId?: string
  time?: Date
  tideHeight?: number
  surgeHeight?: number
  ornamentalLevel?: string
}

export type RealtimeTide = {
  locationId?: string
  time?: Date
  tideHeight?: number
}