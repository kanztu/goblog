package main

const (
	cataHtml          = `<li class="%s"><a href="%s">%s</a></li>`
	blog_preview_html = `
	<div class="col-md-12">
		<div class="blog-entry d-md-flex">
			<a href="%s" class="img img-2"
				style="background-image: url(%s);"></a>
			<div class="text text-2 pl-md-4">
				<h3 class="mb-2"><a href="%s">%s</a></h3>
				<div class="meta-wrap">
					<p class="meta">
						<span><i class="icon-calendar mr-2"></i>%s</span>
						<span><a onclick="searchBlogByTag(%d)"><i
									class="icon-folder-o mr-2"></i>%s</a></span>
					</p>
				</div>
				<p class="mb-4">%s</p>
				<p><a href="%s" class="btn-custom">Read More <span
							class="ion-ios-arrow-forward"></span></a></p>
			</div>
		</div>
	</div>
	`
	tag_cloud_html = `<a onclick="fetchTagResult(%d)" class="tag-cloud-link">%s</a>`
)
