// Components
import App from './App.vue'
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'

export class Core {

    public init(): void {
      const app = createApp(App)
      registerPlugins(app)
      app.mount('#app')
    }

}
