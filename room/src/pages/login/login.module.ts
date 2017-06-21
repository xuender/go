import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { LoginPage } from './login';
import { RoomProvider } from '../../providers/room/room';

@NgModule({
	declarations: [
		LoginPage,
	],
	imports: [
		IonicPageModule.forChild(LoginPage),
	],
	exports: [
		LoginPage
	],
	providers: [
		RoomProvider
	]
})
export class LoginPageModule {}
