@php
  $x = 1;
  $y = $x + 1;
@endphp

@if($x < $y)
  @php echo 'x is less than y'; @endphp
@endif

@foreach(range(1, 10) as $i)
  @continue($i % 2 === 0)
  @php echo $i; @endphp
@endforeach
