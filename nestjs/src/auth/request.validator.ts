import {Request} from "express";
import {verify} from 'jsonwebtoken';
import {UnauthorizedException} from "@nestjs/common";
import {UserContext} from "./auth.context";

/**
 * verifies if request has valid token, if so userContext is added to request
 * @throws UnauthorizedException
 */
export const validateRequest = (request: Request): Request => {
    const bearerHeader = request.headers.authorization;
    const accessToken = bearerHeader && bearerHeader.split(' ')[1];
    try {
        request.userContext = <UserContext>(verify(accessToken, process.env.JWT_SECRET));
        return request;
    } catch (ex) {
        throw new UnauthorizedException(ex.message);
    }
}