import { EntityManager } from '../entityManager.service';

export abstract class Repository {
  protected constructor(readonly entityManager: EntityManager) {}
}
