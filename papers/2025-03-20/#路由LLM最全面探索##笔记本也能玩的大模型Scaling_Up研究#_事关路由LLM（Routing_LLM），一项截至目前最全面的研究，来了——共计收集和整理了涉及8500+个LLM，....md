# #路由LLM最全面探索##笔记本也能玩的大模型Scaling Up研究# 事关路由LLM（Routing LLM），一项截至目前最全面的研究，来了——共计收集和整理了涉及8500+个LLM，...

**URL**: https://weibo.com/6105753431/Pjgszzzde

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B7%AF%E7%94%B1LLM%E6%9C%80%E5%85%A8%E9%9D%A2%E6%8E%A2%E7%B4%A2%23&amp;extparam=%23%E8%B7%AF%E7%94%B1LLM%E6%9C%80%E5%85%A8%E9%9D%A2%E6%8E%A2%E7%B4%A2%23" data-hide=""><span class="surl-text">#路由LLM最全面探索#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%AC%94%E8%AE%B0%E6%9C%AC%E4%B9%9F%E8%83%BD%E7%8E%A9%E7%9A%84%E5%A4%A7%E6%A8%A1%E5%9E%8BScaling+Up%E7%A0%94%E7%A9%B6%23&amp;extparam=%23%E7%AC%94%E8%AE%B0%E6%9C%AC%E4%B9%9F%E8%83%BD%E7%8E%A9%E7%9A%84%E5%A4%A7%E6%A8%A1%E5%9E%8BScaling+Up%E7%A0%94%E7%A9%B6%23" data-hide=""><span class="surl-text">#笔记本也能玩的大模型Scaling Up研究#</span></a> <br><br>事关路由LLM（Routing LLM），一项截至目前最全面的研究，来了——<br><br>共计收集和整理了涉及8500+个LLM，在12个Benchmark上的共2亿条性能记录！<br><br>先来简单科普一下路由LLM。<br><br>这种方法主要是把像ChatGPT、Qwen、DeepSeek这些成型的LLM当作 “专家” ，当给一个输入的时候，有分类能力的Router（路由器）就会把这个输入分配给合适的LLM处理。<br><br>如此一来，就能实现高性能、低计算消耗、低幻觉等目标。<br><br>而来自中山大学和普渡大学的研究人员在基于上述海量的记录做了一番探索之后，发现了一个现象，叫做Model-level Scaling Up。<br><br>一言蔽之，就是一个好的Router，可以让路由LLM范式的性能随着LLM候选数量的增加迅速变强。<br><br>随后，他们通过这些数据构建了针对Router设计的评测RouterEval。<br><br>值得注意的是，其他研究人员，也可以通过RouterEval在很少的计算资源下（如笔记本、单卡GPU上）就能参与到该路由LLM的研究当中。<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2FOQFUcemTEmGC0eKUO_Fuiw" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1hzm99rc631j30u0091dit.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1hzm9aggvavj311w0juwlw.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

中山大学和普渡大学的研究团队进行了一项关于路由LLM（Routing LLM）的全面研究，涉及8500多个大型语言模型（LLM）在12个基准测试上的2亿条性能记录。路由LLM通过将输入分配给合适的LLM（如ChatGPT、Qwen、DeepSeek）来处理，以实现高性能、低计算消耗和减少幻觉。研究发现，随着LLM候选数量的增加，路由LLM的性能显著提升，称为“Model-level Scaling Up”。团队还开发了RouterEval评测工具，使其他研究人员能在有限计算资源（如笔记本、单卡GPU）下参与路由LLM研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T02:31:10Z
- **目录日期**: 2025-03-20
