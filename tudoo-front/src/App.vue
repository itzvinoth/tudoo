<template>
	<div class="container">
		<div id="nav">
			<router-link to="/tasks">List</router-link>|
			<router-link to="/about">Scheduled</router-link>|
			<router-link to="/today">Today</router-link>
		</div>
		<router-view />
	</div>	
</template>

<style src="./assets/styles/app.scss" lang="scss"></style>
