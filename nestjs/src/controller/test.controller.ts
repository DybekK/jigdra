import {Controller, Get, Req, UseGuards} from "@nestjs/common";
import {AuthGuard} from "../auth/auth.guard";
import {Request} from "express";

@Controller("api")
export class TestController {
    @Get("test")
    @UseGuards(AuthGuard)
    public getTest(@Req() req: Request | any) {
        return req.userContext;
    }
}