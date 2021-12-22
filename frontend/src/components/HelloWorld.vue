<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 mb-3">
          Bem vindo ao WNotificator
        </h1>

        <p class="subheading font-weight-regular">
          Configure sua cidade e receba notificações personalizadas sobre o clima no seu WhatsApp.
          <br>
        </p>
        <v-btn
          color="indigo darken-1 white--text"
          depressed
          elevation="6"
          x-large
        >Criar Conta</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  export default {
    name: 'HelloWorld',

    data: () => ({
    }),
  }
</script>
