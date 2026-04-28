@extends('layouts.app')

@section('content')
  <page-header>
    <h1>{{ $title }}</h1>
  </page-header>

  <content-grid>
    @foreach($posts as $post)
      <post-card>
        <img slot="image" src="{{ $post->thumbnail }}" alt="{{ $post->title }}">
        <h2 slot="title">{{ $post->title }}</h2>
        <p slot="excerpt">{{ $post->excerpt }}</p>
        <a slot="link" href="{{ $post->url }}">Read more</a>
      </post-card>
    @endforeach
  </content-grid>

  @if($pagination)
    <page-navigation current="{{ $page }}" total="{{ $totalPages }}">
    </page-navigation>
  @endif

  <site-footer year="{{ date('Y') }}">
    <nav-menu slot="links" orientation="vertical">
      @foreach($footerLinks as $link)
        <nav-item href="{{ $link->url }}">{{ $link->label }}</nav-item>
      @endforeach
    </nav-menu>
  </site-footer>
@endsection
