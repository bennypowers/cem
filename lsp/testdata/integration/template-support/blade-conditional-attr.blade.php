@extends('layouts.app')

@section('content')
  <site-header theme="{{ $theme_mode }}">
    <nav-menu orientation="horizontal">
      @foreach($menu_items as $item)
      <nav-item href="{{ $item->url }}">{{ $item->label }}</nav-item>
      @endforeach
    </nav-menu>
  </site-header>

  @foreach($docs as $doc)
    <rh-tile id="tile-{{ $doc->tagName }}" compact bleed @if($doc->comingSoon) disabled @endif>
      <img slot="image" src="{{ $doc->screenshotPath }}" alt="{{ $doc->tagName }}">
      <h3 slot="headline"><a href="{{ $doc->overviewHref }}">{{ $doc->tagName }}</a></h3>
      @if($doc->comingSoon)
        <div slot="footer">Coming Soon</div>
      @endif
    </rh-tile>
  @endforeach
@endsection
