<script lang="ts">
	let props = $props<{
		videoId: string;
		showModal: boolean;
		setShowModal: (value: boolean) => void;
	}>();
</script>

<div class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
	<button
		class="absolute top-4 right-4 text-white cursor-pointer"
		onclick={() => props.setShowModal(false)}
	>
		close
	</button>
	<div class="bg-white p-6 rounded-lg shadow-lg">
		<video controls autoplay class="w-full h-auto">
			<source src={`/api/video/${props.videoId}`} type="video/mp4" />
			<track
				kind="captions"
				src={`/api/video/${props.videoId}/captions.vtt`}
				srclang="en"
				label="English"
			/>
		</video>
	</div>
</div>
