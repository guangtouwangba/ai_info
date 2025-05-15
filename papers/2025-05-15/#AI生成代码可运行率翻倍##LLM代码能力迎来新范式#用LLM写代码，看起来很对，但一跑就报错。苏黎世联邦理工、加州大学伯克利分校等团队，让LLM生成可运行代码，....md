# #AI生成代码可运行率翻倍##LLM代码能力迎来新范式#用LLM写代码，看起来很对，但一跑就报错。苏黎世联邦理工、加州大学伯克利分校等团队，让LLM生成可运行代码，...

**URL**: https://weibo.com/6105753431/PrWGyCnDH

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E7%94%9F%E6%88%90%E4%BB%A3%E7%A0%81%E5%8F%AF%E8%BF%90%E8%A1%8C%E7%8E%87%E7%BF%BB%E5%80%8D%23&amp;extparam=%23AI%E7%94%9F%E6%88%90%E4%BB%A3%E7%A0%81%E5%8F%AF%E8%BF%90%E8%A1%8C%E7%8E%87%E7%BF%BB%E5%80%8D%23" data-hide=""><span class="surl-text">#AI生成代码可运行率翻倍#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23LLM%E4%BB%A3%E7%A0%81%E8%83%BD%E5%8A%9B%E8%BF%8E%E6%9D%A5%E6%96%B0%E8%8C%83%E5%BC%8F%23&amp;extparam=%23LLM%E4%BB%A3%E7%A0%81%E8%83%BD%E5%8A%9B%E8%BF%8E%E6%9D%A5%E6%96%B0%E8%8C%83%E5%BC%8F%23" data-hide=""><span class="surl-text">#LLM代码能力迎来新范式#</span></a><br><br>用LLM写代码，看起来很对，但一跑就报错。<br><br>苏黎世联邦理工、加州大学伯克利分校等团队，让LLM生成可运行代码，通过率提高了50%！<br><br>他们做的事很简单——在原有的“预测下一个token”机制外，加了一层“类型守门员”。  <br><br>也就是在生成代码的过程中，实时检查它的类型是否正确，从源头上避免生成出的代码无法编译问题。<br><br>具体来说，他们有3大创新：<br><br>1、前缀自动机（Prefix Automaton）：它能动态记录代码上下文中的类型状态，比如一个变量是啥类型、函数期望返回啥类型等。比起传统语法检查，它更像一个“实时类型分析器”。<br><br>2、类型约束解码（Type-Constrained Decoding）：每当模型要生成一个新 token，先用一个前缀自动机看这个token合不合类型规则，不符合就直接砍掉这个路径，换一个。<br><br>3、类型可达性算法：用来判断“当前这段半成品代码，有没有可能写成目标类型的表达式”，从而避免模型进入“写不下去”的死胡同。<br><br>为验证方法的可行性，团队选择了TypeScript做实验，一方面是因为它的类型系统足够有代表性，另一方面也有大量真实项目可以参考。<br><br>模型在两个经典数据集（HumanEval 和 MBPP）上测试，结果很亮眼：<br><br>- 编译失败率减少超过50%；<br>- 功能正确率提升3.5%~5.5%；<br>- 在修bug任务上，准确率提升37%。<br><br>更重要的是，这种方式对模型体积几乎没有要求，无论是LLaMA家族还是其他开源 LLM，只要支持token-level解码逻辑，就可以集成。<br><br>此外，论文还强调这套机制不仅限于代码生成，还可用在代码翻译（比如把 Python转成TypeScript）、修复（自动Debug），甚至能用在文档生成任务中。<br><br>见此情形，不少开发者”脑洞大开“，纷纷献策：<br><br>- 有人提出可以用这种方式训练专精单一语言的“垂类LLM”；<br><br>- 也有人建议将这种类型约束机制和现有的LSP工具（语言服务器）结合，进一步提升智能；<br><br>- 还有开发者呼吁扩展到像Rust、Haskell、甚至Coq这种更强类型语言，把编译器的静态检查能力真正用起来。<br><br>未来的大模型或许真的会像程序员一样，一边写一边检查，一边Debug，生成的代码直接就能跑。<br><br>项目代码已开源：eth-sri/type-constrained-code-generation<br>论文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fabs%2F2504.09246" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1g9follxcj30zk0wmnb8.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1g9fpqxtqj30v011gqn7.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

苏黎世联邦理工和加州大学伯克利分校团队提出新方法，通过引入"类型守门员"机制显著提升LLM生成代码的可运行率。该技术包含三大创新：前缀自动机动态追踪类型状态、类型约束解码过滤非法token、类型可达性算法避免死胡同。在TypeScript测试中，编译失败率降低超50%，功能正确率提升3.5%-5.5%，修bug准确率提高37%。该方法兼容各类LLM，未来可扩展至代码翻译、修复等场景，并有望应用于Rust等强类型语言。项目代码已开源，标志着AI代码生成进入"实时类型检查"新范式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T09:03:13Z
- **目录日期**: 2025-05-15
