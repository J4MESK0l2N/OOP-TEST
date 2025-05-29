export class User {
  constructor(
    public id: number,
    public username: string,
    public password: string,
    public role: string = `user`
  ) {}

  isAdmin(): boolean {
    return this.role === "admin";
  }

  PublicUserToJson() {
    return {
      id: this.id,
      username: this.username,
      role: this.role,
    };
  }
}
