// package: door
// file: door.proto

import * as jspb from "google-protobuf";

export class Door extends jspb.Message {
  getPath(): string;
  setPath(value: string): void;

  getMethod(): MethodEnum;
  setMethod(value: MethodEnum): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Door.AsObject;
  static toObject(includeInstance: boolean, msg: Door): Door.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Door, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Door;
  static deserializeBinaryFromReader(message: Door, reader: jspb.BinaryReader): Door;
}

export namespace Door {
  export type AsObject = {
    path: string,
    method: MethodEnum,
    data: Uint8Array | string,
  }
}

export enum MethodEnum {
  GET = 0,
  POST = 1,
  PUT = 2,
  DELETE = 3,
}

