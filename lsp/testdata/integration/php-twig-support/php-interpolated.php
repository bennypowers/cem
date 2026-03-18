<?php
$brand_color = get_theme_mod('brand_color', '#0073aa');
$card_variant = get_option('card_variant', 'outlined');
?>
<style>
:root {
  --brand-color: <?php echo esc_attr($brand_color); ?>;
}
</style>
<product-card variant="<?php echo esc_attr($card_variant); ?>">
  <img slot="image" src="<?php echo esc_url($image); ?>" alt="">
  <span slot="price"><?php echo esc_html($price); ?></span>
</product-card>
<?php if ($show_sidebar): ?>
<sidebar-widget collapsible>
  <?php dynamic_sidebar('main-sidebar'); ?>
</sidebar-widget>
<?php endif; ?>
