<template>
  <v-app>
      <v-app-bar
        color="deep-purple"
        dark
      >
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

        <div class="d-flex align-center">
          <v-img
            alt="Vuetify Name"
            class="shrink mt-1 hidden-sm-and-down"
            contain
            min-width="100"
            src="./assets/logo_name.png"
          />
        </div>
      </v-app-bar>

      <v-navigation-drawer
        v-model="drawer"
        absolute
        temporary
      >
        <v-list
          nav
          dense
        >
          <v-list-item-group
            v-model="group"
            active-class="deep-purple--text text--accent-4"
          >
            <v-list-item>
              <v-list-item-icon>
                <v-icon>mdi-home</v-icon>
              </v-list-item-icon>
              <v-list-item-title>Home</v-list-item-title>
            </v-list-item>

            <v-list-item>
              <v-list-item-icon>
                <v-icon>mdi-account</v-icon>
              </v-list-item-icon>
              <v-list-item-title>Account</v-list-item-title>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </v-navigation-drawer>

    <v-main>
      <HelloWorld/>
      <wether-card initialCity="San Francisco"/>
      <Footer/>
    </v-main>
  </v-app>
</template>

<script>
import Footer from './components/Footer.vue';
import HelloWorld from './components/HelloWorld';
import WetherCard from './components/WetherCard';

export default {
  name: 'App',

  components: {
    HelloWorld,
    WetherCard,
    Footer
  },

  data: () => ({
      drawer: false,
      group: null,
  }),
};
</script>
