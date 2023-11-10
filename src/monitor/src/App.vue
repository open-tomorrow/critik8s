<script lang="ts" setup>
  import axios from 'axios'
  import VueAxios from 'vue-axios'
  import { AMQPService } from '@/amqp'
  import NodePanel from './components/MainPanel.vue'
  import CriticPanel from './components/CriticPanel.vue'

  const links = [
    'Dashboard',
  ]
</script>

<script lang="ts">
  export default {
    data: () => ({
      links: [],
      nodes: []
    }),

    methods: {
    subscribeNodes() {

      axios.get('http://localhost/api/auth').then((response) => {

      const {
        rabbitmqPath,
        rabbitmqUsername,
        rabbitmqPassword,
        dataCollectorRoutingKey
      } = response.data;

      let amqpService = new AMQPService("localhost/"+rabbitmqPath,
                                        "80",
                                        rabbitmqUsername,
                                        rabbitmqPassword,
                                        "logs_topic",
                                        dataCollectorRoutingKey)
      amqpService.subscribe()
      amqpService.messagesSubject.subscribe({
        next: (v) => (this.nodes as any[]).push(v),
      });
    })

    }
  },
  }
</script>

<template>
  <v-app id="inspire">
    <v-app-bar flat>
      <v-container class="mx-auto d-flex align-center justify-center">

        <v-btn
          v-for="link in links"
          :key="link"
          :text="link"
          variant="text"
        ></v-btn>

        <v-spacer></v-spacer>

      </v-container>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col cols="2">
            <v-sheet rounded="lg">
              <v-list rounded="lg">

                <v-list-item title="RWO" color="" link></v-list-item>

                <v-divider class="my-2"></v-divider>

                <v-list-item @click="subscribeNodes"
                  color="grey-lighten-4"
                  link
                  title="Test Sub"
                ></v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet min-height="50vh" rounded="lg">
              <span>{{ nodes }}</span>
              <node-panel/>
            </v-sheet>
            <v-divider></v-divider>
            <v-sheet min-height="30vh" rounded="lg">
              <!--  -->
              <critic-panel/>
            </v-sheet>
          </v-col>

        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>
