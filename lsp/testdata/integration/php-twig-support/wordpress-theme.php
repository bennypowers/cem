<?php
/**
 * Template Name: Full Width
 * Description: A WordPress page template using web components
 */
$site_name = get_bloginfo('name');
$show_posts = have_posts();
?>
<!DOCTYPE html>
<html <?php language_attributes(); ?>>
<head><title><?php echo esc_html($site_name); ?></title></head>
<body <?php body_class(); ?>>
  <site-header theme="dark">
    <h1 style="color: var(--color-text)"><?php echo esc_html($site_name); ?></h1>
  </site-header>

  <main>
    <?php
    if ($show_posts) :
      while (have_posts()) : the_post();
    ?>
    <blog-card>
      <h2 slot="heading"><?php the_title(); ?></h2>
      <?php the_content(); ?>
    </blog-card>
    <?php
      endwhile;
    endif;
    ?>
  </main>

  <site-footer year="<?php echo date('Y'); ?>">
    <p>&copy; <?php echo esc_html($site_name); ?></p>
  </site-footer>
</body>
</html>
