import {Request} from "express";
import {sign} from 'jsonwebtoken';
import {validateRequest} from "../../../src/auth/request.validator";
import {config} from 'dotenv';
import {UnauthorizedException} from "@nestjs/common";

// initialize env variables
config();

const generateToken = (payload: any): string => sign(payload, process.env.JWT_SECRET, {expiresIn: 30 * 60});

describe("request validator test", () => {
    it("should return request object with userContext if token is valid", () => {
        // given
        const payload = {email: "test"};
        const mockedRequest = {
            headers: {
                authorization: " " + generateToken(payload)
            }
        } as Request

        // when
        const actual = validateRequest(mockedRequest).userContext;

        // then
        expect(actual.email).toBe(payload.email);
    });

    it("should throw UnauthorizedException if token is not provided", () => {
        // given
        const mockedRequest = {
            headers: {
                authorization: ""
            }
        } as Request

        expect(
            // when
            () => validateRequest(mockedRequest)
            // then
        ).toThrow(UnauthorizedException);
    });

    it("should throw UnauthorizedException if token is invalid", () => {
        // given
        const mockedRequest = {
            headers: {
                authorization: "invalid token"
            }
        } as Request

        expect(
            // when
            () => validateRequest(mockedRequest)
            // then
        ).toThrow(UnauthorizedException);
    });
});