import {Column, Model, Table, PrimaryKey, IsUUID} from "sequelize-typescript";

@Table
export class WorkspaceUser extends Model {
    @IsUUID(4)
    @PrimaryKey
    @Column
    id: string

    @Column
    nickname: string

    @Column({ defaultValue: true })
    isActive: boolean;
}