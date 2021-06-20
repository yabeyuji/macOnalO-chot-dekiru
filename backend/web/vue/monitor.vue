<template>
	<b-container fluid class="text-light text-center">
		<b-row>
			<b-col
				cols="12"
				class="btn-secondary border_color_black"
				:class="[commonClass]"
			>
				<b-row>
					<b-col
						cols="5"
						class="btn-primary border_color_gray"
						:class="[commonClass]"
					>
					reserve
					</b-col>
				</b-row>
				<b-row>
					<b-col
						cols="2"
						v-for="(reserve, index) in reserves"
					>
						<b-row>
							<b-col
								cols="12"
								class="btn-info border_color_gray"
								:class="[commonClass]"
							>
								{{ reserve }}
							</b-col>
						</b-row>
					</b-col>
				</b-row>
			</b-col>
		</b-row>
		<!-- ---------------------------- -->
		<b-row>
			<b-col
				cols="12"
				class="btn-secondary border_color_black"
				:class="[commonClass]"
			>
				<b-row>
					<b-col
						cols="5"
						class="btn-primary border_color_gray"
						:class="[commonClass]"
					>
					assemble
					</b-col>
				</b-row>
				<b-row>
					<b-col
						cols="2"
						v-for="(assemble, index) in assembles"
					>
						<b-row>
							<b-col
								cols="12"
								class="btn-info border_color_gray"
								:class="[commonClass]"
							>
								{{ assemble }}
							</b-col>
						</b-row>
					</b-col>
				</b-row>
			</b-col>
		</b-row>
		<!-- ---------------------------- -->
		<b-row>
			<b-col
				cols="12"
				class="btn-secondary border_color_black"
				:class="[commonClass]"
			>
				<b-row>
					<b-col
						cols="5"
						class="btn-primary border_color_gray"
						:class="[commonClass]"
					>
					complete
					</b-col>
				</b-row>
				<b-row>
					<b-col
						cols="2"
						v-for="(complete, index) in completes"
					>
						<b-row>
							<b-col
								cols="12"
								class="btn-info border_color_gray"
								:class="[commonClass]"
							>
								{{ complete }}
							</b-col>
						</b-row>
					</b-col>
				</b-row>
			</b-col>
		</b-row>

		<!-- ---------------------------- -->
		<b-row>
			<b-col
				cols="12"
				class="btn-secondary border_color_black"
				:class="[commonClass]"
			>
				<b-row>
					<b-col
						cols="5"
						class="btn-primary border_color_gray"
						:class="[commonClass]"
					>
					pass
					</b-col>
				</b-row>
				<b-row>
					<b-col
						cols="2"
						v-for="(pass, index) in passes"
					>
						<b-row>
							<b-col
								cols="12"
								class="btn-info border_color_gray"
								:class="[commonClass]"
							>
								{{ pass }}
							</b-col>
						</b-row>
					</b-col>
				</b-row>
			</b-col>
		</b-row>




	</b-container>
</template>

<script>
module.exports = {
	props: {
		'commonClass': {type: String, default: ''},
		'reserves': { type: Array, default: [] },
		'assembles': { type: Array, default: [] },
		'completes': { type: Array, default: [] },
		'passes': { type: Array, default: [] },
	},

	methods: {
	},

	components: {
	},


};
</script>
