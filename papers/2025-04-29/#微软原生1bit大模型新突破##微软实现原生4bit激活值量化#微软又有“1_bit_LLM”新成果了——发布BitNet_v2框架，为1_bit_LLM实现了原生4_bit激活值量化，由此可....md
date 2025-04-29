# #微软原生1bit大模型新突破##微软实现原生4bit激活值量化#微软又有“1 bit LLM”新成果了——发布BitNet v2框架，为1 bit LLM实现了原生4 bit激活值量化，由此可...

**URL**: https://weibo.com/6105753431/PpwdNcvtN

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BE%AE%E8%BD%AF%E5%8E%9F%E7%94%9F1bit%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%96%B0%E7%AA%81%E7%A0%B4%23&amp;extparam=%23%E5%BE%AE%E8%BD%AF%E5%8E%9F%E7%94%9F1bit%E5%A4%A7%E6%A8%A1%E5%9E%8B%E6%96%B0%E7%AA%81%E7%A0%B4%23" data-hide=""><span class="surl-text">#微软原生1bit大模型新突破#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BE%AE%E8%BD%AF%E5%AE%9E%E7%8E%B0%E5%8E%9F%E7%94%9F4bit%E6%BF%80%E6%B4%BB%E5%80%BC%E9%87%8F%E5%8C%96%23&amp;extparam=%23%E5%BE%AE%E8%BD%AF%E5%AE%9E%E7%8E%B0%E5%8E%9F%E7%94%9F4bit%E6%BF%80%E6%B4%BB%E5%80%BC%E9%87%8F%E5%8C%96%23" data-hide=""><span class="surl-text">#微软实现原生4bit激活值量化#</span></a><br><br>微软又有“1 bit LLM”新成果了——<br><br>发布BitNet v2框架，为1 bit LLM实现了原生4 bit激活值量化，由此可充分利用新一代GPU（如GB200）对4 bit计算的原生支持能力。同时减少内存带宽&amp;提升计算效率。<br><br>之前，微软持续研究BitNet b1.58 <a href="https://weibo.com/6105753431/PohITnfFD" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a> ，把LLM的权重量化到1.58-bit，显著降低延迟、内存占用等推理成本。<br><br>然鹅BitNet b1.58激活值还是8-bit，这就导致没办法充分利用新一代硬件的4 bit计算能力，计算环节出现效率瓶颈。<br><br>而BitNet v2引入了新的H-BitLinear模块。这个模块在激活量化前加入Hadamard变换，成功把尖锐的激活分布揉成了更接近高斯分布的形态，异常值大幅减少，从而真正实现了4bit激活。详情请见文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405160820635205867" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">微软1bit LLM新研究：4bit激活值，可充分利用新一代GPU进行4bit计算</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0xt2yl6a5j30rs0fmgnq.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

微软推出BitNet v2框架，实现1bit大模型的原生4bit激活值量化，充分利用新一代GPU（如GB200）对4bit计算的支持，提升计算效率并减少内存占用。此前BitNet b1.58已将权重量化至1.58bit，但激活值仍为8bit，导致计算效率受限。BitNet v2引入H-BitLinear模块，通过Hadamard变换优化激活分布，减少异常值，成功实现4bit激活，显著提升硬件利用率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T22:03:51Z
- **目录日期**: 2025-04-29
