<script lang="ts">
  let count: number = $state(0)
  const increment = () => {
    count += 1
  }
</script>

<button type="button" class="btn" onclick={increment}>
  Count is {count}
</button>
<button class="btn">Default</button>
