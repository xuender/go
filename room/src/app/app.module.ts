import { BrowserModule } from '@angular/platform-browser';
import { ErrorHandler, NgModule } from '@angular/core';
import { IonicApp, IonicErrorHandler, IonicModule } from 'ionic-angular';

import { MyApp } from './app.component';
import { HomePage } from '../pages/home/home';
import { ListPage } from '../pages/list/list';

import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { LoginPageModule } from '../pages/login/login.module';
import { RoomProvider } from '../providers/room/room';

@NgModule({
	declarations: [
		MyApp,
		HomePage,
		ListPage
	],
	imports: [
		BrowserModule,
		IonicModule.forRoot(MyApp),
		LoginPageModule,
	],
	bootstrap: [IonicApp],
		entryComponents: [
		MyApp,
		HomePage,
		ListPage
	],
	providers: [
		StatusBar,
		SplashScreen,
		{provide: ErrorHandler, useClass: IonicErrorHandler},
		RoomProvider
	]
})
export class AppModule {}
