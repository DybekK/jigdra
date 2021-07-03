export interface DatabaseConfiguration {
    host: string,
    port: number,
    username: string,
    password: string,
    name: string
}

export interface AppConfiguration {
    database: DatabaseConfiguration
}
