import {UserContext} from "../../src/modules/auth/auth.context";

declare global {
    namespace Express {
        interface Request {
            userContext?: UserContext
        }
    }
}