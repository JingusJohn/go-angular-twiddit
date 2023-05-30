import { Component } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  email: string = "";
  username: string = "";
  password: string = "";
  confirmPassword: string = "";

  constructor(private authService: AuthService) { }

  login() {
    console.log("login");
    
    this.authService.login(this.email, this.username, this.password, this.confirmPassword).then((result) => {
      console.log("result: ", result);
    });
  }
}
