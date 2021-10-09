<template>
	<div id="nav">
		<router-link to="/">List</router-link>|
		<router-link to="/about">Scheduled</router-link>|
		<router-link to="/today">Today</router-link>
	</div>
	<router-view />
</template>

<style src="./assets/styles/app.scss" lang="scss"></style>
