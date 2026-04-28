{{-- Quoted interpolations in attributes (supported) --}}
<my-element id="el-{{ $id }}"
            variant="{{ $variant }}"
            title="{!! $rawTitle !!}"
            data-count="{{ $items->count() }}">
  <span slot="label">{{ $label }}</span>
</my-element>

{{-- Multiple interpolations in one attribute --}}
<status-badge class="badge badge-{{ $type }} {{ $active ? 'active' : '' }}">
</status-badge>
