// package: 
// file: door.proto

import * as jspb from "google-protobuf";

export class Door extends jspb.Message {
  getType(): number;
  setType(value: number): void;

  getSuccess(): boolean;
  setSuccess(value: boolean): void;

  getError(): string;
  setError(value: string): void;

  hasUser(): boolean;
  clearUser(): void;
  getUser(): User | undefined;
  setUser(value?: User): void;

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
    type: number,
    success: boolean,
    error: string,
    user?: User.AsObject,
  }
}

export class User extends jspb.Message {
  getNick(): string;
  setNick(value: string): void;

  getToken(): string;
  setToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    nick: string,
    token: string,
  }
}

