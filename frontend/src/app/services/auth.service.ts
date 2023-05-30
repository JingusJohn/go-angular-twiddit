import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  baseUrl: string;

  constructor() {
    // get from env in the future, for now hardcode
    this.baseUrl = "http://localhost:3000";
  }

  public async isAuthenticated(): Promise<boolean> {
    let result = await fetch(`${this.baseUrl}/api/protected/auth/health`);
    let json = await result.json();
    console.log(json);

    return false;
  }

  public async login(email: string, username: string, password: string, confirmPassword: string): Promise<any> {
    let result = await fetch(`${this.baseUrl}/api/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email,
        username,
        password,
        confirm_password: confirmPassword
      }),
    });
    let json = await result.json();
    return json;
  }
}
