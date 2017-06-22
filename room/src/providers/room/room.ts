import { Injectable } from '@angular/core';
import 'rxjs/add/operator/map';
import { $WebSocket } from 'angular2-websocket/angular2-websocket'
import { Door } from '../../generated/door_pb'

@Injectable()
export class RoomProvider {
	public ws: $WebSocket;
	constructor() {
		this.ws = new $WebSocket('ws://localhost:8888/ws');
		this.ws.onMessage((msg: MessageEvent) => {
			console.log('onMessage', msg.data);
			const reader = new FileReader();
			reader.readAsArrayBuffer(msg.data);
			reader.onload = function (e) {
				const buf = new Uint8Array(reader.result);
				const d = Door.deserializeBinary(buf);
				console.log('door', d);
				const w = d.toObject();
				console.log('door', w);
			}
		}, {autoApply: false});
		console.log('Hello RoomProvider Provider');
	}
	send() {
		this.ws.send("xxx");
	}
}
