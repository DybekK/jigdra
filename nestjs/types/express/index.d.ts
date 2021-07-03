import {UserContext} from "../../src/auth/auth.context";

declare global {
    namespace Express {
        interface Request {
            userContext?: UserContext
        }
    }
}