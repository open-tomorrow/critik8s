/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import vuetify from './vuetify'

// Types
import type { App } from 'vue'
import * as Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

export function registerPlugins (app: App) {
  app
    .use(vuetify)
    .use(VueAxios, axios)
}
