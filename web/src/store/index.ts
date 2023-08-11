import { createPinia } from 'pinia';
import { createPersistedState } from 'pinia-plugin-persistedstate';

const store = createPinia();
store.use(createPersistedState());

export { store };

export * from './menu';
export * from './user';
export * from './openim_session';
export * from './openim_message';

export default store;
