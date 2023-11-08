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

                <v-list-item
                  color="grey-lighten-4"
                  link
                  title="New"
                ></v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col>
            <v-sheet min-height="50vh" rounded="lg">
              <!--  -->
              Main window
            </v-sheet>
            <v-divider></v-divider>
            <v-sheet min-height="30vh" rounded="lg">
              <!--  -->
              Critics
            </v-sheet>
          </v-col>

        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
  const links = [
    'Dashboard',
  ]
</script>

<script>
  export default {
    data: () => ({
      links: [
        'Dashboard',
      ],
    }),
  }
</script>
