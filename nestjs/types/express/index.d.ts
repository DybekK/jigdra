import { UserContext } from '../../src/modules/auth/guard/auth.context';

declare global {
  namespace Express {
    interface Request {
      userContext?: UserContext;
    }
  }
}
