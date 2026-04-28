@extends('layouts.app')

@section('content')
  <page-layout>
    @foreach($sections as $section)
      <content-section id="section-{{ $section->id }}">
        @if($section->hasHeader)
          <section-header variant="{{ $section->theme }}">
            @foreach($section->items as $item)
              <card-element
                title="{{ $item->title }}"
                href="{{ $item->url }}"
                @if($item->featured) featured @endif>
                @if($item->image)
                  <img slot="media" src="{{ $item->image }}" alt="{{ $item->title }}">
                @endif
                <span slot="description">{{ $item->description }}</span>
              </card-element>
            @endforeach
          </section-header>
        @endif
      </content-section>
    @endforeach
  </page-layout>
@endsection
