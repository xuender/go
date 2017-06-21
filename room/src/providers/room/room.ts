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
			let a=new Uint8Array(msg.data.parts);
			console.log('door a', a);
			const d = Door.deserializeBinary(a);
			console.log('door', d);
			console.log('door', d.toObject());
		}, {autoApply: false});
		console.log('Hello RoomProvider Provider');
	}
	send() {
		this.ws.send("xxx");
	}
}
