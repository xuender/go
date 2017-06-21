import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { HomePage } from '../home/home';
import { RoomProvider } from '../../providers/room/room';

@IonicPage()
@Component({
	selector: 'page-login',
	templateUrl: 'login.html',
})
export class LoginPage {
	constructor(
		public navCtrl: NavController,
		public navParams: NavParams,
		public room: RoomProvider
	) {
	}
	ionViewDidLoad() {
		console.log('ionViewDidLoad LoginPage');
		// 判断是否已经登录
		// this.navCtrl.setRoot(HomePage);
	}
	login() {
		this.room.send();
		this.navCtrl.setRoot(HomePage);
	}
}
