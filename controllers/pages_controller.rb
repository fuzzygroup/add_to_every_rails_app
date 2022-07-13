class PagesController < ApplicationController
  
  def index
    content = read_markdown("index.md")
    @content = markdown(content)
  end

  def about
    content = read_markdown("about.md")
    @content = markdown(content)
  end  
  
  def contact
    content = read_markdown("contact.md")
    @content = markdown(content)
  end  
  
  def dedication
    content = read_markdown("dedication.md")
    @content = markdown(content)
  end
  
  def faq
    content = read_markdown("faq.md")
    @content = markdown(content)
  end

  def mission
    content = read_markdown("mission.md")
    @content = markdown(content)
  end

  def values
    content = read_markdown("values.md")
    @content = markdown(content)
  end
  
  private
  
  def markdown(text)
    options = %i[
      hard_wrap autolink no_intra_emphasis tables fenced_code_blocks
      disable_indented_code_blocks strikethrough lax_spacing space_after_headers
      quote footnotes highlight underline
    ]
    Markdown.new(text, *options).to_html.html_safe
  end
  
  def read_markdown(filename)
    File.read(File.join(Rails.root, 'app/views/pages', filename))
  end
end
