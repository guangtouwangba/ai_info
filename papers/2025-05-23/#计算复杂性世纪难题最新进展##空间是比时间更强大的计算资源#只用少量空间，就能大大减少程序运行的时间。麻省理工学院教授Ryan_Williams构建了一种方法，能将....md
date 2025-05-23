# #计算复杂性世纪难题最新进展##空间是比时间更强大的计算资源#只用少量空间，就能大大减少程序运行的时间。麻省理工学院教授Ryan Williams构建了一种方法，能将...

**URL**: https://weibo.com/6105753431/Pt12R5zgb

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%AE%A1%E7%AE%97%E5%A4%8D%E6%9D%82%E6%80%A7%E4%B8%96%E7%BA%AA%E9%9A%BE%E9%A2%98%E6%9C%80%E6%96%B0%E8%BF%9B%E5%B1%95%23&amp;extparam=%23%E8%AE%A1%E7%AE%97%E5%A4%8D%E6%9D%82%E6%80%A7%E4%B8%96%E7%BA%AA%E9%9A%BE%E9%A2%98%E6%9C%80%E6%96%B0%E8%BF%9B%E5%B1%95%23" data-hide=""><span class="surl-text">#计算复杂性世纪难题最新进展#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%A9%BA%E9%97%B4%E6%98%AF%E6%AF%94%E6%97%B6%E9%97%B4%E6%9B%B4%E5%BC%BA%E5%A4%A7%E7%9A%84%E8%AE%A1%E7%AE%97%E8%B5%84%E6%BA%90%23&amp;extparam=%23%E7%A9%BA%E9%97%B4%E6%98%AF%E6%AF%94%E6%97%B6%E9%97%B4%E6%9B%B4%E5%BC%BA%E5%A4%A7%E7%9A%84%E8%AE%A1%E7%AE%97%E8%B5%84%E6%BA%90%23" data-hide=""><span class="surl-text">#空间是比时间更强大的计算资源#</span></a><br><br>只用少量空间，就能大大减少程序运行的时间。<br><br>麻省理工学院教授Ryan Williams构建了一种方法，能将任何算法转化为一种占用空间大幅缩减的形式。<br><br>在计算复杂性理论中，有两个关键的概念：包含所有能在合理时间内解决的问题的“P类”以及与之对应的空间复杂度类别“PSPACE类”。<br><br>这两类的关系是复杂性理论的核心问题之一。所有P类问题都属于PSPACE类，因为快速算法根本没有足够时间占满计算机的内存空间。<br><br>但如果逆命题也成立，两类就将等价——空间与时间将具备相当的计算能力<br><br>然而复杂性理论家们认为，空间是远比时间强大的计算资源，因为算法可以反复利用同一小块存储空间，而时间却无法回溯。<br><br>如何证明这个直觉性的判断？要证实PSPACE类大于P类，必须证明某些计算无法在有限时间内完成。<br><br>或者反过来：给定特定空间预算的算法，能够解决那些仅需稍大时间预算的算法所能处理的所有问题。<br><br>1975年，John Hopcroft、Leslie Valiant与Wolfgang Paul运用了一种称为“模拟”的方法——通过转化现有算法，生成能解决相同问题但时空配比不同的新算法。<br><br>三人设计出了一种通用模拟程序，总能实现空间压缩：无论输入何种算法，输出的等效算法其空间预算总会略小于原算法的时间预算。<br><br>然而此后研究陷入停滞，一直到五十年后。<br><br>Williams这一突破性研究的起点，源于树评估问题：当算法的内存空间被严格限制时，能否完成对特定树结构数据的计算？<br><br>2023年，James Cook与研究者Ian Mertz设计的算法以远低于预期的空间需求解决了树评估问题，这种新算法揭示：数据竟可压缩叠放。<br><br>Williams深入研究后灵光乍现：Cook-Mertz方法本质上是通用型空间压缩工具，何不借此构建新版时空通用模拟？<br><br>基于“可压缩卵石”的模拟，使新算法的空间需求急剧下降，约等于原算法时间预算的平方根。<br><br>Williams通过这个模拟证明了关于空间计算能力的肯定性结论：占用较少空间的算法，能够解决那些需要消耗更多时间的问题。<br><br>随后他将其转化为关于时间计算能力的否定性结论：至少存在某些问题，若不投入比空间更多的时间资源就无解。<br><br>不过，Williams的第二结论看似解决了“P对PSPACE”难题。但却存在着微观与宏观的本质尺度差异。要最终证明PSPACE大于P，研究者仍需将这个差距扩大数个数量级。<br><br>感兴趣的朋友可以阅读全文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.quantamagazine.org%2Ffor-algorithms-a-little-memory-outweighs-a-lot-of-time-20250521%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1oeef98v2j31rx17tx6p.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

麻省理工学院Ryan Williams教授在计算复杂性理论取得突破性进展，提出了一种将算法转化为空间占用大幅缩减形式的新方法。研究通过"可压缩卵石"模拟技术，证明空间资源比时间更强大——少量空间可显著减少运行时间。该成果改进了1975年Hopcroft等人的空间压缩技术，使新算法空间需求降至原时间预算的平方根级别。虽然尚未完全解决"P对PSPACE"这一世纪难题，但为证明PSPACE类大于P类提供了重要理论依据，揭示了空间作为计算资源的独特优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T00:03:01Z
- **目录日期**: 2025-05-23
