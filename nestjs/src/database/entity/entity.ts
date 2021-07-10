import {Column, Entity, Index, ManyToOne, OneToMany, PrimaryGeneratedColumn} from "typeorm";

@Entity()
export class WorkspaceUser {
    @PrimaryGeneratedColumn("uuid")
    id: string;

    @Index({unique: true})
    @Column()
    userId: string

    @Column()
    nickname: string;

    @Column({default: true})
    isActive: boolean;

    @ManyToOne(() => Workspace, workspace => workspace.workspaceUsers)
    workspace: Promise<Workspace>;
}

@Entity()
export class Workspace {
    @PrimaryGeneratedColumn("uuid")
    id: string;

    @Column()
    name: string

    @OneToMany(() => WorkspaceUser, workspaceUser => workspaceUser.workspace)
    workspaceUsers: Promise<WorkspaceUser[]>;
}