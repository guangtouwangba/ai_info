# #怎样用AI写代码更省钱##氛围编程是什么#Vibe Coding（氛围编程），说的是你用自然语言提需求，AI写代码，哪里不对，再用自然语言告诉它，AI再改，直到你满意为...

**URL**: https://weibo.com/6105753431/Ps4my2FxM

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%80%8E%E6%A0%B7%E7%94%A8AI%E5%86%99%E4%BB%A3%E7%A0%81%E6%9B%B4%E7%9C%81%E9%92%B1%23&amp;extparam=%23%E6%80%8E%E6%A0%B7%E7%94%A8AI%E5%86%99%E4%BB%A3%E7%A0%81%E6%9B%B4%E7%9C%81%E9%92%B1%23" data-hide=""><span class="surl-text">#怎样用AI写代码更省钱#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B0%9B%E5%9B%B4%E7%BC%96%E7%A8%8B%E6%98%AF%E4%BB%80%E4%B9%88%23&amp;extparam=%23%E6%B0%9B%E5%9B%B4%E7%BC%96%E7%A8%8B%E6%98%AF%E4%BB%80%E4%B9%88%23" data-hide=""><span class="surl-text">#氛围编程是什么#</span></a><br><br>Vibe Coding（氛围编程），说的是你用自然语言提需求，AI写代码，哪里不对，再用自然语言告诉它，AI再改，直到你满意为止。<br><br>Fred Benenson分享了Vibe Coding的优劣势：纯用AI生成代码，感觉太爽了，但耗费的token太多、太贵了。<br><br>具体而言，Fred表示氛围编程的以下特点：<br><br>1、用AI写代码，开销惊人<br><br>他用Claude 3.7 Sonne写代码，仅4月Token费用就超$260，累计已超$1000t。主要原因是，每次你改一点，Claude得把之前所有的代码重新发一遍，token开销太大了。<br><br>2、AI代码不仅啰嗦，还不能跑<br><br>他举了个例子：2018年手写的minimax算法是400行，而Claude写的是627行。Claude的代码注释、结构都挺完备，但整体架构混乱，逻辑冗余，而且还跑不了。<br><br>这是因为，Claude会为了“显得聪明”去生成很多辅助函数，结果一点也不优雅，维护成本还更高。<br><br>3、究其原因，是“token经济”的反向激励<br><br>Claude和很多AI助手按token收费，token越多公司赚得越多，这就意味着：<br><br>- Claude写得越长、改得越频繁，公司赚得越多<br><br>- Claude并没有动力优化代码结构，因为冗余才赚钱<br><br>- 你的每一次对话，都加大了上下文体积，越聊越贵<br><br>作者把这叫做“Perverse Incentive”（反向激励），也就是产品目标（写好代码）和商业模型（多赚token费用）存在天然冲突。<br><br>“Claude不会因为写得烂而惩罚自己，只会让你多花点钱。”<br><br>4、简洁代码反而更容易出错<br><br>此前，Giskard发布的Phare benchmark表明：当你让模型“尽量简洁”时，它更容易出现幻觉（hallucination），也就是胡说八道。【图3】<br><br>对此，作者提供了几种自救策略：<br><br>1. 强制让AI先写计划：别一上来就让AI写代码，先让它列清结构，效果会更好。<br><br>2. 内置提示加权限控制：在Claude的内置提示里，写明“必须经过我允许才能生成代码”，让AI别自作主张写。虽然Claude有时候还是不听，但反复提醒能减少无用输出。<br><br>3. 用上Git该删就删：Claude的一堆代码做好备份，该删就删，能省不少token费用。<br><br>4. 用小模型省钱省心：Claude 3.5 Haiku每token只要Claude 3.7的四分之一，而且更简洁，很多任务够用。<br><br>Fred Benenson调侃说：这背后的商业决策值得反思，AI生成代码看似高效，但其实把你引入了token陷阱，越用越乱，越乱越贵。<br><br>感兴趣的小伙伴可以点击原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Ffredbenenson.medium.com%2Fthe-perverse-incentives-of-vibe-coding-23efbaf75aee" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1h7awu4h4j32vs29mqv6.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1h7ax8hhij312w0ongre.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1h7azxr69j30lz0zk7d3.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

这篇微博讨论了"氛围编程"(Vibe Coding)的优缺点。使用AI(如Claude)通过自然语言交互生成代码虽然便捷，但存在三大问题：1)费用高昂，因每次修改都需重发全部代码；2)生成的代码冗长且常无法运行，维护成本高；3)AI按token收费的商业模式导致其缺乏优化代码的动力。作者建议通过让AI先写计划、限制自动生成、及时删除冗余代码、使用小模型等方法降低成本。最终指出，当前的AI编程模式可能陷入"越用越乱，越乱越贵"的token陷阱。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T07:02:56Z
- **目录日期**: 2025-05-16
