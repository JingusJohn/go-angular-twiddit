import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { FeedComponent } from './components/feed/feed.component';
import { SignupComponent } from './components/signup/signup.component';

const routes: Routes = [
  { path: "", component: FeedComponent },
  { path: "login", component: LoginComponent },
  { path: "feed", component: FeedComponent },
  { path: "signup", component: SignupComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
