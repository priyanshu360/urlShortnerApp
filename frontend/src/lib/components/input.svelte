<script>
  let inputUrl = '';
  let shortenedUrl = '';


  async function shortenURL() {
    try {
      // Make your API call here
      const response = await fetch('http://localhost:8081/create-short-url', {
        method: 'POST',
        // headers: {
        //   'Content-Type': 'application/json',
        // },
        body: JSON.stringify({ "long_url": inputUrl }),
      });

      if (response.ok) {
        const data = await response.json();
        shortenedUrl = data.data;
      } else {
        console.error('API call failed');
      }
    } catch (error) {
      console.error('An error occurred while making the API call', error);
    }
  }
</script>

<div class="flex items-center justify-center bg-gray-100" style="height: 90vh;">
  <div class="bg-white shadow-lg rounded-lg p-8 w-96">
    <input type="text" placeholder="Enter a long URL" bind:value={inputUrl} class="block w-full p-3 border rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
    <button on:click={shortenURL} class="mt-4 block w-full px-4 py-3 bg-blue-500 hover:bg-blue-700 text-white rounded-r-lg">Shorten</button>
    {#if shortenedUrl}
      <p class="mt-4">Shortened URL: <a href={shortenedUrl} target="_blank">{shortenedUrl}</a></p>
    {/if}
  </div>
</div>
