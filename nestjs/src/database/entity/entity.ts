import {Column, Entity, Index, ManyToMany, ManyToOne, OneToMany, PrimaryGeneratedColumn} from "typeorm";
import {worker} from "cluster";

export abstract class Model {
    @PrimaryGeneratedColumn("uuid")
    id: string;
}

@Entity()
export class WorkspaceUser extends Model {
    @Index({unique: true})
    @Column()
    userId: string

    @Column()
    nickname: string;

    @Column({default: true})
    isActive: boolean;

    @ManyToOne(() => Workspace, workspace => workspace.workspaceUsers, {
        cascade: true
    })
    workspace: Promise<Workspace>;

    @ManyToMany(() => Task, task => task.workspaceUsers, {
        cascade: true
    })
    tasks: Promise<Task[]>;
}

@Entity()
export class Workspace extends Model {
    @Column()
    name: string

    @OneToMany(() => WorkspaceUser, workspaceUser => workspaceUser.workspace, {
        cascade: true
    })
    workspaceUsers: Promise<WorkspaceUser[]>;

    @OneToMany(() => Task, task => task.workspace, {
        cascade: true
    })
    tasks: Promise<Task[]>;
}

export class Task {
    @Column()
    title: string;

    @Column()
    description: string;

    @ManyToOne(() => Workspace, workspace => workspace.tasks, {
        cascade: true
    })
    workspace: Promise<Workspace>;

    @ManyToMany(() => WorkspaceUser, workspaceUser => workspaceUser.tasks, {
        cascade: true
    })
    workspaceUsers: Promise<WorkspaceUser>
}