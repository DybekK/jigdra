import {Column, Entity, Index, ManyToOne, OneToMany, PrimaryGeneratedColumn} from "typeorm";
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

    @ManyToOne(() => Workspace, workspace => workspace.workspaceUsers)
    workspace: Promise<Workspace>;

    @OneToMany(() => Task, task => task.workspace)
    tasks: Promise<Task[]>;
}

@Entity()
export class Workspace extends Model {
    @Column()
    name: string

    @OneToMany(() => WorkspaceUser, workspaceUser => workspaceUser.workspace)
    workspaceUsers: Promise<WorkspaceUser[]>;

    @OneToMany(() => Task, task => task.workspace)
    tasks: Promise<Task[]>;
}

export class Task {
    @Column()
    title: string;

    @Column()
    description: string;

    @ManyToOne(() => Workspace, workspace => workspace.tasks)
    workspace: Promise<Workspace>;

    @ManyToOne(() => WorkspaceUser, workspaceUser => workspaceUser.tasks)
    workspaceUser: Promise<WorkspaceUser>
}