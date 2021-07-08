import { EntityManager } from './entityManager.service';

export abstract class PrismaRepository {
  protected constructor(readonly em: EntityManager) {}
}

